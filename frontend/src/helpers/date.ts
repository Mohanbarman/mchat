export const parseTime = (date: Date): string => {
    return date.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
    });
};

export const parseDate = (date: Date): string => {
    return date.toLocaleDateString("en-US", {
        day: "2-digit",
        month: "2-digit",
        year: "2-digit",
    });
};

export const parseFriendlyDate = (date: Date): string => {
    switch (getDayDiff(date)) {
        case 1:
            return "Today";
        case 2:
            return "Yesterday";
        default:
            return parseDate(date);
    }
};

export const getDayDiff = (date: Date): number => {
    return new Date(Date.now() - date.getTime()).getUTCDate();
};
