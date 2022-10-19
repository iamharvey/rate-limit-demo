import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 150,
    duration: '30s',
    thresholds: {
        // the rate of successful checks should be < 50%
        checks: ['rate>0.3'],
    },
};

export default function () {
    const res = http.get('http://127.0.0.1:4000/');
    check(res, {
        'status code 429': (r) => r.status == 429,
    });
    sleep(1)
}