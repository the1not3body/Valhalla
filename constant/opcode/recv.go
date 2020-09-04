package opcode

const (
	RecvLoginRequest               byte = 0x01
	RecvLoginChannelSelect         byte = 0x04
	RecvLoginWorldSelect           byte = 0x05
	RecvLoginCheckLogin            byte = 0x08
	RecvLoginCreateCharacter       byte = 0x09
	RecvLoginSelectCharacter       byte = 0x0B
	RecvChannelPlayerLoad          byte = 0x0C
	RecvLoginNameCheck             byte = 0x0D
	RecvLoginNewCharacter          byte = 0x0E
	RecvLoginDeleteChar            byte = 0x0F
	RecvPing                       byte = 0x12
	RecvReturnToLoginScreen        byte = 0x14
	RecvChannelUserPortal          byte = 0x17
	RecvCHannelChangeChannel       byte = 0x18
	RecvChannelEnterCashShop       byte = 0x19
	RecvChannelPlayerMovement      byte = 0x1A
	RecvChannelPlayerStand         byte = 0x1B
	RecvChannelPlayerUseChair      byte = 0x1C
	RecvChannelMeleeSkill          byte = 0x1D
	RecvChannelRangedSkill         byte = 0x1E
	RecvChannelMagicSkill          byte = 0x1F
	RecvChannelDmgRecv             byte = 0x21
	RecvChannelPlayerSendAllChat   byte = 0x22
	RecvChannelEmote               byte = 0x23
	RecvChannelNpcDialogue         byte = 0x27
	RecvChannelNpcDialogueContinue byte = 0x28
	RecvChannelNpcShop             byte = 0x29
	RecvChannelInvMoveItem         byte = 0x2D
	RecvChannelAddStatPoint        byte = 0x36
	RecvChannelPassiveRegen        byte = 0x37
	RecvChannelAddSkillPoint       byte = 0x38
	RecvChannelSpecialSkill        byte = 0x39
	RecvChannelCharacterInfo       byte = 0x3F
	RecvChannelLieDetectorResult   byte = 0x45
	RecvChannelCharacterReport     byte = 0x49
	RecvChannelSlashCommands       byte = 0x4C
	RecvChannelCharacterUIWindow   byte = 0x4E
	RecvChannelPartyInfo           byte = 0x4F
	RecvChannelGuildManagement     byte = 0x51
	RecvChannelGuildReject         byte = 0x52
	RecvChannelAddBuddy            byte = 0x55
	RecvChannelUseMysticDoor       byte = 0x58
	RecvChannelMobControl          byte = 0x6A
	RecvChannelMobEffect           byte = 0x6B
	RecvChannelNpcMovement         byte = 0x6F
)
