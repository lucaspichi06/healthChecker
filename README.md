# healthChecker

``` javascript
/**
The HealthMonitor class is a singleton providing a central place to register
resources that will be checked periodically for their health.
The registration is done via monitor().
The status/health of all registered resources is obtained via the check() method.
*/

// Example usage:
/*
healthMonitor.monitor({
type: 'serviceUrl',
name: 'graphql',
handle: this.uri,
critical: true
});
status = await healthMonitor.check()
*/

const _ = require('lodash');
class HealthMonitor {
// the following functions are stubbed for brevity
// the real implementation consist of resource specific code that checks the resource health
// eg. checkRedisClient connects to redis and executes a trivial qiery.
static async checkRedisClient(resourceName, redisClient) {
return Promise.resolve({resource: resourceName, status: 'ok'});
}
static async checkPostgresPromiseClient(resourceName, postgresClient) {
return Promise.resolve({resource: resourceName, status: 'ok'});
}

    static async checkElasticClient(resourceName, elasticClient) {
        return Promise.resolve({resource: resourceName, status: 'ok'});
    }

    static async checkServiceUrl(resourceName, url) {
        return Promise.resolve({resource: resourceName, status: 'ok'});
    }

    constructor() {
        if (!HealthMonitor.instance) {
            this.monitors = new Map();
            this.criticalResources = new Set();
            HealthMonitor.instance = this;
        }
        return HealthMonitor.instance;
    }

    monitor(resource) {
        const { type, name, handle, critical } = resource;
        if (!type || !name || !handle) {
            return;
        }
        let resourceByType = this.monitors.get(type);
        if (!resourceByType) {
            resourceByType = new Map();
        }
        resourceByType.set(name, handle);
        this.monitors.set(type, resourceByType);
        if (critical) {
            this.criticalResources.add(name);
        }
    }

    async check() {
        const healthPromises = [];
        const handlerFuncs = {
            redisClient: HealthMonitor.checkRedisClient,
            elasticsearchClient: HealthMonitor.checkElasticClient,
            serviceUrl: HealthMonitor.checkServiceUrl,
            postgresClient: HealthMonitor.checkPostgresClient,
            postgresPromiseClient: HealthMonitor.checkPostgresPromiseClient,
            postgresPool: HealthMonitor.checkPostgresPool
        };
        for (const pair of this.monitors) {
            const checkFunc = handlerFuncs[pair[0]];
            if (checkFunc) {
                for (const resources of pair[1]) {
                    healthPromises.push(
                        checkFunc(resources[0], resources[1])
                    );
                }
            }
        }
        const timeout = this.maxHealthcheckDuration;
        const runPromises = [
            Promise.all(healthPromises).then(status => Promise.reject(status)),
            new Promise((resolve, reject) => {
                setTimeout(() => {
                    reject('timeout');
                }, timeout);
            })
        ];
        let results = [];
        try {
            await Promise.all(runPromises);
        } catch (err) {
            if (err === 'timeout') {
                return {
                    status: 503,
                    error: 'timed out while checking resources',
                    failed: Array.from(this.criticalResources.values())
                };
            }
            results = err;
        }
        const failed = [];
        const state = {
            status: 200
        };
        for (const r of results) {
            if (r.status !== 'ok') {
                if (this.criticalResources.has(r.resource)) {
                    state.status = 503;
                }
                failed.push(r.resource);
            }
            _.set(state, r.resource, _.omit(r, 'resource'));
        }
        if (failed.length > 0) {
            state.failed = failed;
        }
        return state;
    }
}
const instance = new HealthMonitor();
Object.freeze(instance);

module.exports = instance;
```