import { UserAddModal } from "../../components";
import { useCreateConversation } from "../../http/conversation/conversation.hooks";
import { commonActions } from "../../redux/common";
import { actions } from "../../redux/conversations/conversationSlice";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";

export const AddUser: React.FC = () => {
    const { isAddUserModalOpen } = useAppSelector((s) => s.common);
    const dispatch = useAppDispatch();
    const { execute, error } = useCreateConversation();

    const onClose = () => dispatch(commonActions.closeAddUserModal());

    const onSubmit = async (email: string) => {
        const { success } = await execute({ email });
        if (!success) return;

        onClose();
        dispatch(actions.add(success.data.data));
        dispatch(actions.setActive(success.data.data.id));
    };

    return <UserAddModal isOpen={isAddUserModalOpen} onClose={onClose} onSubmit={onSubmit} error={error} />;
};
