import {AggregateType} from '../jestTestType';
import {deepCopy, filterData, addImageClass, hasShowPassword, conversionRgba, hasShowTooltip} from '@/utils/index.ts';

describe('test-method', () => {
    // deepcopy
    it('test-success-Array-deepcopy', () => {
        const copyData1: AggregateType['DeepCopyType']['Data'][] = [{name: 1}, {name: 2}];
        const resultData1: AggregateType['DeepCopyType']['Data'][] = [{name: 1}];
        expect(deepCopy(copyData1)).not.toEqual(deepCopy(resultData1));
    }),
    it('test-success-Object-deepcopy', () => {
        const copyObj1 = {name: 1, info: {age: 2, sex: '男'}};
        expect(deepCopy(copyObj1)).not.toBe(deepCopy(copyObj1));
        expect(deepCopy(copyObj1)).toEqual(copyObj1);
    }),
    it ('test-null-deepcopy', () => {
        expect(deepCopy(null)).toBeNull();
    }),
    it ('test-undefined-deepcopy', () => {
        expect(deepCopy(undefined)).toBeUndefined();
    }),
    it('test-string-deepcopy', () => {
        expect(deepCopy('string')).toMatch(/string/);
    })

    // filterData
    it('test-success-filterData', () => {
        const data = [{name: 1}, {name: 1}];
        expect(filterData(data, 'name')).toEqual(filterData(data));
        expect(filterData(data, 'name')).toEqual([{name: 1}]);
    }),
    it('test-[]-filterData', () => {
        expect(filterData([])).toEqual([]);
    }),

    // addImageClass
    it('test-success-addImageClass', () => {
        const data = [{name: 1}];
        const moreData =  Array(10).fill(true);
        expect(addImageClass(data, true)).toBe('img-show');
        expect(addImageClass(moreData, true)).toBe('');
    }),
    it('test-[]-addImageClass', () => {
        expect(addImageClass([], true)).toBe('img-show');
        expect(addImageClass([], false)).toBe('img-empty');
    }),

    // hasShowPassword
    it('test-success-hasShowPassword', () => {
        expect(hasShowPassword(false, '12312')).toBe('*****');
        expect(hasShowPassword(true, '132132')).toBe('132132');
    }),
    it('test-empty-hasShowPassword', () => {
        expect(hasShowPassword()).toBeUndefined();
    }),

    // conversionRgba
    it('test-success-conversionRgba', () => {
        expect(conversionRgba('#ffffff', '.6')).toBe('rgba(255, 255,255, .6)');
    }),
    it('test-error-conversionRgba', () => {
        expect(conversionRgba('#fff', '.6')).toBeNull();  
        expect(conversionRgba('', '.6')).toBeNull();
        expect(conversionRgba('#ffffff', '')).toBe('rgba(255, 255,255, )')
        expect(conversionRgba()).toBeNull();  
    }),
    it('test-hasShowTooltip', () => {
        expect(hasShowTooltip(false, 'adasdsad'));
    })
});
