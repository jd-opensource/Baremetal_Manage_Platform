type DetailType = {detail: {[x: string]: string;};};

type CurrencyType = {[x: string]: string};

interface TipStatusType {
    tooltipStatus: {
        [x: string]: {
            showTooltip: boolean;
        }
    }
};

interface ShowPasswordType {
    internetPort1: boolean;
    internetPort2: boolean;
    outLoginPassword: boolean;
    info: string;
};

export {
    DetailType,
    CurrencyType,
    TipStatusType,
    ShowPasswordType
};
