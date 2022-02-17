import { UseFormSetError } from "react-hook-form";

export const handleFieldErrors = (
    res: Record<string, any>,
    callback: UseFormSetError<any>
) => {
    for (const key in res) {
        callback(key as any, {
            message: res[key],
            type: "manual",
        });
    }
};
