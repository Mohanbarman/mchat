package conversation

type CreateDTO struct {
	UserID string `mapstructure:"user_id"`
}

type SendMessageDTO struct {
	UserID string `mapstructure:"user_id"`
	Text   string `mapstructure:"text"`
	FileID string `mapstructure:"file_id"`
}

type ReadConversationDTO struct {
	ConversationID string `mapstructure:"conversation_id"`
}
