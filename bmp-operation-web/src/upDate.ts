/**
 * @file
 * @author
*/
import {ElMessageBox} from 'element-plus';

class UpdateOpt {
    flag: number = 0;
    hasLoop: number = 0;
    timer: null | number = null;
    env: string = process.env.VUE_APP_ENV;

    constructor() {
        this.init();
        this.handleVisibility();
    }

    handleVisibility = () => {
        document.addEventListener('visibilitychange', () => {
            const status = {
                'hidden': () => {
                    clearInterval(this.timer as number);
                    this.hasLoop = 0;
                },
                'visible': () => this.init()
            }
            status[document.visibilityState]();
        })
    }

    init = () => {
        if (!this.hasLoop) {
            this.loopReq();
        }
        // 定时执行，自动更新逻辑(每10s检测一次)
        this.timer = setInterval(async () => {
            this.loopReq();
        }, 1000 * 10);
    }

    loopReq = async () => {
        this.hasLoop = 1;
        // 本地
        const localVersion: string = this.getLocalHash();
        // 线上
        const newVersion = await this.getPreHash();
        // 如果不一样，就进行刷新
        if(localVersion !== newVersion && ![newVersion, localVersion].includes('')) {
            if (this.flag > 0 || Object.is(this.env, 'dev')) return;
            ElMessageBox.confirm(
                '检测到有新版本，是否更新？',
                '版本提示',
                {
                    confirmButtonText: '更新',
                    cancelButtonText: '取消',
                    type: 'warning'
                }
            )
            .then(() => {
                this.flag = 0;
                window.location.reload();
            })
            .catch(() => {
                this.flag = 0;
            })
            this.flag ++;
        }
    }

    getAppHash = (scripts: string | any[] | HTMLCollectionOf<HTMLScriptElement>) => {
        let localVersion = '';
        for(let i = 0; i < scripts.length; i++) {
            let src = scripts[i].getAttribute("src");
            if (src && src.indexOf("app.") !== -1) {
                // 正则返回中间版本号(如果没有，返回空)
                const regRes = /app\.(.*?).js/;
                if(regRes.test(src)) {
                    const matchResult = regRes.exec(src);
                    if (matchResult !== null) {
                        localVersion = matchResult[1];
                    }
                    else {
                        // 处理未匹配到的情况，可能是赋一个默认值，或者抛出异常
                        localVersion = '未找到匹配的版本号';
                        throw new Error(localVersion);
                    }
                }
            }
        }
        return localVersion;
    }

    // 获取本地的app.js版本号
    getLocalHash = () => {
        return this.getAppHash(document.getElementsByTagName('script'));
    }

    // 获取线上的app.js版本号
    getPreHash = () => {
        return new Promise((resolve, reject) => {
            // 加上时间戳，防止缓存 
            fetch('/?_time=' + Date.now()).then(async res => {
                // 转成字符串判断
                const html = await res.text();
                const doc = new DOMParser().parseFromString(html, 'text/html');
                const newVersion = this.getAppHash(doc.getElementsByTagName('script'));
                resolve(newVersion);
            })
            .catch(err => {
                reject(err);
                throw new Error('获取版本号失败');
            })
        })
    }
}

new UpdateOpt();