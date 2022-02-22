package conversation

type CreateDTO struct {
	UserID string `mapstructure:"user_id" validation:"required"`
}

type SendMessageDTO struct {
	UserID string `mapstructure:"user_id" validation:"required"`
	Text   string `mapstructure:"text" validation:"required"`
	FileID string `mapstructure:"file_id"`
}

type ReadConversationDTO struct {
	ConversationID string `mapstructure:"conversation_id" validation:"required"`
}

type TypingEventDTO struct {
	ConversationID string `mapstructure:"id" validation:"required"`
}
