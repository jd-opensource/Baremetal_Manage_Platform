interface RuleFormType {
    userName: string;
    resourceType: string;
    idcId: string;
    instanceId: string;
    dimension: string;
    monitoringIndicators: string[];
    // hours: string;
    // value1: string;
    keyData: any;
};

export {
    RuleFormType
}