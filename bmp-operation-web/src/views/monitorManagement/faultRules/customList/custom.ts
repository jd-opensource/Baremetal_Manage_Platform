/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import useCustom from 'hooks/custom/useCustom.ts';

const name: string = methodsTotal.humpConversion('alertRuleList');

const customOperate = useCustom(name).info;

export default customOperate;