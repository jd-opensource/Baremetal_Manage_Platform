/**
 * @file
 * @author
*/

class AlarmHistoryOpt {
    alaramHistoryDiaLog: Ref<boolean> = ref<boolean>(false)
    ruleId: Ref<string> = ref<string>('');

    alarmHistoryClick = (ruleId: string) => {
        this.ruleId.value = ruleId;
        this.alaramHistoryDiaLog.value = !this.alaramHistoryDiaLog.value;
    }

    alarmHistoryCancel = (type: boolean) => {
        this.alaramHistoryDiaLog.value = type;
    }
}

export default AlarmHistoryOpt;
