/**
 * @file
 * @author
*/
import {language} from 'utils/publicClass.ts';

class RulesOpt {
    rules = reactive({
        userName: [
            {
                trigger: 'blur',
                required: true,
                message: language.t('allAlarmRulesList.search.errorTip.userName')
            }
        ]
    })
}

export default RulesOpt;
