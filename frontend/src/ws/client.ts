export interface ISendMessagePayload {
    text: string;
    userID: string;
}

export class WsClient {
    private socket: WebSocket;

    constructor(ws: WebSocket) {
        this.socket = ws;
    }

    sendMessage(payload: ISendMessagePayload) {
        const data = {
            action: "conversation/send",
            payload: {
                user_id: payload.userID,
                text: payload.text,
            },
        };
        this.sendJson(data);
    }

    readConversation(payload: string) {
        const data = {
            action: "conversation/read",
            payload: {
                conversation_id: payload,
            },
        };
        this.sendJson(data);
    }

    login(token: string) {
        const data = {
            action: "auth/login",
            payload: {
                token,
            },
        };
        this.sendJson(data);
    }

    sendJson(data: Record<string, any>) {
        this.socket.send(JSON.stringify(data));
    }
}
