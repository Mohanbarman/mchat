export interface IMessageProps {
    type: "text" | "image";
    text: string;
    isMe: boolean;
    time: Date;
    state: "sent" | "seen" | "delivered";
    id: string;
    onClick: (conversationId: string) => any;
}
