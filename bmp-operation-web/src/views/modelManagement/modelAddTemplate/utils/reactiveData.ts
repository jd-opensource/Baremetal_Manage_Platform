import {StoreGeneric, storeToRefs} from 'pinia'; // store
import {ReactiveDataType} from '../typeManagement';
import store from 'store/index.ts';

interface DataType {
    modelFormStore: StoreGeneric;
    reactiveData: ReactiveDataType;
};

const reactiveData = () => {
    const data: DataType = {
        modelFormStore: store.modeFormInfo(),
        reactiveData: reactive<ReactiveDataType>({
            ...storeToRefs(store.modeFormInfo()).data.value
        })
    };

    return {
        ...data
    };
};

export default reactiveData;
