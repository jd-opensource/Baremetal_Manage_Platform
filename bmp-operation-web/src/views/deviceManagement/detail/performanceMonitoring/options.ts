/**
 * @file
 * @author
*/

// import echarts from './utils/publicQuote';
import * as echarts from 'echarts';

const kb: number = 1000;
const mb: number = kb * 1000;
const gb: number = mb * 1000;
const tb: number = gb * 1000;
const pb: number = tb * 1000;
const eb: number = pb * 1000;

const reg = (str: string) => str.replace(/\.00$/, '');
let newMyChart: any;
const publicOption = (legendData: string[], ...args: any) => {
    const status = Object.is(args[5], 'bigSize');
    return new Promise((resolve) => {
        resolve({
            title: {
                top: '10px',
                textStyle: {
                    fontSize: 14
                }
            },
            tooltip: {
                trigger: 'axis',
                formatter: function (params: {seriesName: string; value: number; color: string; axisValue: string;}[]) {
                    let res: string = '';
                    for (let i = 0, l = params.length; i < l; i++) {
                        const yearSecond = new Date(args[6][i] * 1000);
                        const yearSecondZero = new Date(args[6][0] * 1000);
                        const monthDateTime = params[i].axisValue.replace(/\n/g, ',');
                        const seconds = isNaN(yearSecond.getSeconds()) ? yearSecondZero.getSeconds() < 10 ? '0' + yearSecondZero.getSeconds() : yearSecondZero.getSeconds() : yearSecond.getSeconds() < 10 ? '0' + yearSecond.getSeconds() : yearSecond.getSeconds();
                        res += `<p>${isNaN(yearSecond.getFullYear()) ? yearSecondZero.getFullYear() : yearSecond.getFullYear()}-${monthDateTime.split(',').reverse()[0] + ' ' + monthDateTime.split(',').reverse()[1] + ':' + seconds}</p>`
                        res += '<span style="display:inline-block;margin-right:5px; border-radius: 50%; width:4px;height:4px;background-color: #fff; border-width: 2px; border-style: solid;' + 'border-color:' + params[i].color + ';"></span>';
                        // res += symbol + ' ' + params[i].seriesName + ' : ' + params[i].value + (args[1] === 'custom' ? flag === '1' ? 'M' : flag === '2' ? 'G' : 'K' : args[4]) +'' + '<br/>';
                        // (args[1] === 'custom' ? convertBytesUnit(Number(params[i].value)) : )
                        res += ' ' + params[i].seriesName + ' : ' + setUnit(args, params[i].value) + '<br/>';
                    }
                    return res;
                },
                textStyle: {
                    color: '#333',
                    fontSize: 12
                }
            },
            legend: {
                data: isMultiDimensional(legendData) ? legendData[0] : legendData,
                left: status ? '-5px' : '13px',
                top: '5%',
                itemWidth: 16,
                itemHeight: 8,
                textStyle: {
                    color: '#333',
                    fontSize: 12
                },
            },
            grid: {
                left: status ? '0' : '4.7%',
                // left: status ? '2.2%' : isMultiDimensional(args[3]) ? args[2].map((item: {data: number[]}) => item.data).flat().some((t: number) => t > 100) ? '4%' : '3%' : args[3].some((t: number) => t > 100) ? '5%' : '4%',
                // left: status ? '2.2%' : isMultiDimensional(args[3]) ? args[2].map((item: {data: number[]}) => item.data).flat().some((t: number) => t > 100) ? '5%' : '1%' : Object.is(args[1], '%') && args[3].some((item: number) => item < 100) ? '5%' : args[3].some((item: number) => item > 100) ? '4%' : '3%',
                // right: '7%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data: isMultiDimensional(args[0]) ? args[0][0].map(function (str: string) {
                    return str.replace(' ', '\n');
                }) : args[0].map(function (str: string) {
                    return str.replace(' ', '\n')
                }),
            },
            yAxis: {
                type: 'value',  
                // min: isMultiDimensional(args[3]) ? 0 : Math.min(...args[3]),
                min: isMultiDimensional(args[3]) ? 0 : args[3].every((item: number) => Number(item) < 0) ? '' : args[3].every((item: number) => Number(item) <= 100) ? 0 : Math.min(...args[3]),
                max: isMultiDimensional(args[3]) ? '' : args[3].every((item: number) => Number(item) < 0) ? '' : args[3].every((item: number) => Number(item) <= 100) ? 100 : '',
                // min: isMultiDimensional(args[3]) ? 0 : Object.is(args[1], '%') && args[3].some((item: number) => item > 100) ? 0 : Math.min(...args[3]),
                // max: isMultiDimensional(args[3]) ? setData(args[3]).some((item: number) => item > 100) ? '' : args[3].some((item: number) => item > 100) ? '' : 100 : Object.is(args[1], '%') && args[2][0]?.data?.length ? 100 : '', 
                axisLabel: {
                    formatter: function(value: number) {
                        return setFormAtter(args, value)
                    }
                }
            },
            dataZoom: [
                {
                    type: 'inside',
                    show: false,
                    start: Math.min(...args[6]),
                    end: Math.max(...args[6]),
                }
            ],
            series: args[2]
        })
    })
};

const addSymbol = (num: string) => {
    return num.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,');
};

const setFormAtter = (args: any, value: number) => {
    const objData = [
        [
            () => ['custom'].includes(args[1]),
            () => {return `${convertBytesUnit(value)}`}
        ],
        [
            () => ['个', '次/秒', '个/秒'].includes(args[1]),
            () => {
                if (value >= 1000) {
                    // .toFixed(2)
                    // parseFloat
                    return reg(addSymbol((String(value / 1000)))) + 'K';
                }
                return `${reg(value.toFixed(2))}`
            }
        ],
        [
            () => args[1] || !args[1],
            () => {return `${reg(value.toFixed(2))}`}
        ]
    ];
    for (const key of objData) {
        if (key[0]()) {
            return key[1]();
        }
    }
}

const setUnit = <T extends number>(args: any, value: T) => {
    if (['custom'].includes(args[1])) {
        return convertBytesUnit(Number(value));
    }
    else if (['个', '次/秒', '个/秒'].includes(args[1])) {
        // if (value >= 1000) {
        //     // .toFixed(2)
        //     // parseFloat
        //     return reg(addSymbol((String(value / 1000)))) + 'K' + args[1];
        // }
        return `${value}${args[1]}`;
    }
    else {
        return reg(value.toFixed(2)) + args[4];
    }
}

const convertBytesUnit = (bytes: number) => {
    const dataList = [
        [
            () => bytes < kb,
            () => {return reg(String(bytes)) + ' B';}
        ],
        [
            () => bytes < mb,
            () => {return reg((bytes / kb).toFixed(2)) + ' K';}
        ],
        [
            () => bytes < gb,
            () => {return reg((bytes / mb).toFixed(2)) + ' M';}
        ],
        [
            () => bytes < tb,
            () => {return reg((bytes / gb).toFixed(2)) + ' G';}
        ],
        [
            () => bytes < pb,
            () => {return reg((bytes / tb).toFixed(2)) + ' T';}
        ],
        [
            () => bytes < eb,
            () => {return reg((bytes / pb).toFixed(2)) + ' P';}
        ]
    ];
    for (const key of dataList) {
        if (key[0]()) {
            return key[1]();
        }
    }
}

const isMultiDimensional = (arr: string[] | number[][]) => {
    return Array.isArray(arr) && arr.some(item => Array.isArray(item));
}

const seriesDataObj = (name: any, ...args: any) => {
    return {
        name,
        type: 'line',
        smooth: true,
        data: args[0],
        itemStyle: {
            color: args[1], // 点的颜色
            lineStyle: {
                color: args[1] // 线的颜色
            }
        },
    };
}

const echartsOpt = (res: any, type: string, myChart: any) => {
    if (!res.yAxis.max) {
        Reflect.deleteProperty(res.yAxis, 'max')
    }
    if (myChart === void 0) {
        myChart = echarts.init(document.getElementById(type));
        newMyChart = myChart;
    }
    res && myChart.setOption(res);
}

export {
    newMyChart,
    publicOption,
    seriesDataObj,
    echartsOpt
};
