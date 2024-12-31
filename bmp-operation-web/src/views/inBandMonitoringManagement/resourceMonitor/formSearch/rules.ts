/**
 * @file
 * @author
*/
import {language} from "utils/publicClass";

class RulesCheckOpt {
    rules = reactive({
        userName: [
            {
                required: true,
                trigger: 'blur',
                message: language.t('resourceMonitor.errorTip.userName')
            }
        ],
        resourceType: [
            {
                required: true,
                trigger: 'change',
                message: language.t('resourceMonitor.errorTip.resourceType')
            }
        ],
        idcId: [
            {
                required: true,
                trigger: 'blur',
                message: language.t('resourceMonitor.errorTip.idcName')
            }
        ],
        instanceId: [
            {
                required: true,
                trigger: 'blur',
                message: language.t('resourceMonitor.errorTip.instanceId')
            }
        ],
        dimension: [
            {
                required: true,
                trigger: 'change',
                message: language.t('resourceMonitor.errorTip.dimensions')
            }
        ],
        monitoringIndicators: [
            {
                required: true,
                trigger: 'blur',
                message: language.t('resourceMonitor.errorTip.monitoringMetrics')
            }
        ]
    });
}

export default RulesCheckOpt;
