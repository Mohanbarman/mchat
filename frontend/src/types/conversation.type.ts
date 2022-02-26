import { IUser } from "./user.type";

export interface IConversation {
    id: string;
    created_at: string;
    updated_at: string;
    is_unread: boolean;
    last_message: string;
    last_message_time: string;
    unread_count: number;
    user: IUser;
    is_typing: boolean;
}
