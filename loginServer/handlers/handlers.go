package handlers

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/Hucaru/Valhalla/common/connection"
	"github.com/Hucaru/Valhalla/common/constants"
	"github.com/Hucaru/Valhalla/common/crypt"
	"github.com/Hucaru/Valhalla/common/packet"
)

// HandlePacket -
func HandlePacket(conn connection.Connection, buffer packet.Packet, isHeader bool) int {
	size := constants.CLIENT_HEADER_SIZE

	if isHeader {
		// Reading encrypted header
		size = crypt.GetPacketLength(buffer)
	} else {
		// Handle data packet
		pos := 0

		opcode := buffer.ReadByte(&pos)

		switch opcode {
		case constants.LOGIN_REQUEST:
			fmt.Println("Login packet received")
			handleLoginRequest(buffer, &pos, conn)
		case constants.LOGIN_CHECK_LOGIN:
			fmt.Println("Check login packet received")
			handleCheckLogin(conn)
		default:
			fmt.Println("UNKNOWN LOGIN PACKET:", buffer)
		}
	}

	return size
}

func handleLoginRequest(p packet.Packet, pos *int, conn connection.Connection) {
	usernameLength := p.ReadInt16(pos)
	username := p.ReadString(pos, int(usernameLength))

	passwordLength := p.ReadInt16(pos)
	password := p.ReadString(pos, int(passwordLength))

	// hash the password
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	var userID uint32
	var user string
	var databasePassword string
	var isLogedIn bool
	var isBanned int
	var isAdmin bool

	err := connection.Db.QueryRow("SELECT userID, username, password, isLogedIn, isBanned, isAdmin FROM users WHERE username=?", username).
		Scan(&userID, &user, &databasePassword, &isLogedIn, &isBanned, &isAdmin)

	result := byte(0x00)

	if err != nil {
		result = 0x05
	} else if hashedPassword != databasePassword {
		result = 0x04
	} else if isLogedIn {
		result = 0x07
	} else if isBanned > 0 {
		result = 0x02
	}

	// -Banned- = 2
	// Deleted or Blocked = 3
	// Invalid Password = 4
	// Not Registered = 5
	// Sys Error = 6
	// Already online = 7
	// System error = 9
	// Too many requests = 10
	// Older than 20 = 11
	// Master cannot login on this IP = 13

	pac := packet.NewPacket()
	pac.WriteByte(constants.LOGIN_RESPONCE)
	pac.WriteByte(result)
	pac.WriteByte(0x00)
	pac.WriteInt32(0)

	if result <= 0x01 {
		pac.WriteUint32(userID)
		pac.WriteByte(0x00)
		if isAdmin {
			pac.WriteByte(0x01)
		} else {
			pac.WriteByte(0x00)
		}
		pac.WriteByte(0x01)
		pac.WriteString(username)
	} else if result == 0x02 {
		// Being banned is not yet implemented
	}

	pac.WriteInt64(0)
	pac.WriteInt64(0)
	pac.WriteInt64(0)
	conn.Write(pac)

	player := conn.GetPlayer()

	player.SetUserID(userID)
	player.SetIsLogedIn(true)

	_, err = connection.Db.Query("UPDATE users set isLogedIn=1 WHERE userID=?", userID)

	if err != nil {
		fmt.Println(err)
	}
}

func handleCheckLogin(conn connection.Connection) {
	// No idea what this packet is for
	pac := packet.NewPacket()
	pac.WriteByte(0x03)
	pac.WriteByte(0x04)
	pac.WriteByte(0x00)
	conn.Write(pac)

	var username string

	userID := conn.GetPlayer().GetUserID()

	err := connection.Db.QueryRow("SELECT username FROM users WHERE userID=?", userID).
		Scan(&username)

	if err != nil {
		fmt.Println("handleCheckLogin database retrieval issue for userID:", userID, err)
	}

	hasher := sha512.New()
	hasher.Write([]byte(username)) // Username should be unique so might as well use this
	hashedUsername := fmt.Sprintf("%x02", hasher.Sum(nil))

	conn.GetPlayer().SetSessionHash(hashedUsername)

	pac = packet.NewPacket()
	pac.WriteByte(constants.LOGIN_SEND_SESSION_HASH)
	pac.WriteString(hashedUsername)
	conn.Write(pac)
}
