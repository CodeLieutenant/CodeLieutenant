async function http<T extends { [x: string]: any }>(url: string, method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE', body?: T, opt?: RequestInit): Promise<Response> {
    if (!opt) {
        opt = {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'X-Requested-With': 'XMLHttpRequest',
            }
        }
    } else {
        opt.method = method;
        opt.headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'X-Requested-With': 'XMLHttpRequest',
            ...opt.headers,
        };
    }


    if ('Content-Type' in opt.headers) {
        switch (opt.headers['Content-Type']) {
            case 'application/json':
                opt.body = JSON.stringify(body);
                break;
            case 'multipart/form-data':
                const data = new FormData();

                for (let item in body) {
                    data.append(item, body[item]);
                }

                opt.body = data;
                break;
        }
    }

    return await fetch(url, opt)
}


export { http };
