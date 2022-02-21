export interface IProps {
    id: string;
    avatar: string;
    name: string;
    message: string;
    messageTime: number;
    isUnread: boolean;
    unreadCount: number;
    onClick: (id: string) => any;
}
