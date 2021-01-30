/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(self["webpackChunktemplate"] = self["webpackChunktemplate"] || []).push([["ts_subscribe_ts"],{

/***/ "./ts/http.ts":
/*!********************!*\
  !*** ./ts/http.ts ***!
  \********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"http\": () => /* binding */ http\n/* harmony export */ });\nvar __assign = (undefined && undefined.__assign) || function () {\n    __assign = Object.assign || function(t) {\n        for (var s, i = 1, n = arguments.length; i < n; i++) {\n            s = arguments[i];\n            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))\n                t[p] = s[p];\n        }\n        return t;\n    };\n    return __assign.apply(this, arguments);\n};\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nvar __generator = (undefined && undefined.__generator) || function (thisArg, body) {\n    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;\n    return g = { next: verb(0), \"throw\": verb(1), \"return\": verb(2) }, typeof Symbol === \"function\" && (g[Symbol.iterator] = function() { return this; }), g;\n    function verb(n) { return function (v) { return step([n, v]); }; }\n    function step(op) {\n        if (f) throw new TypeError(\"Generator is already executing.\");\n        while (_) try {\n            if (f = 1, y && (t = op[0] & 2 ? y[\"return\"] : op[0] ? y[\"throw\"] || ((t = y[\"return\"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;\n            if (y = 0, t) op = [op[0] & 2, t.value];\n            switch (op[0]) {\n                case 0: case 1: t = op; break;\n                case 4: _.label++; return { value: op[1], done: false };\n                case 5: _.label++; y = op[1]; op = [0]; continue;\n                case 7: op = _.ops.pop(); _.trys.pop(); continue;\n                default:\n                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }\n                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }\n                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }\n                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }\n                    if (t[2]) _.ops.pop();\n                    _.trys.pop(); continue;\n            }\n            op = body.call(thisArg, _);\n        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }\n        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };\n    }\n};\nfunction http(url, method, body, opt) {\n    return __awaiter(this, void 0, void 0, function () {\n        var data, item;\n        return __generator(this, function (_a) {\n            switch (_a.label) {\n                case 0:\n                    if (!opt) {\n                        opt = {\n                            method: method,\n                            headers: {\n                                'Content-Type': 'application/json',\n                                'Accept': 'application/json',\n                                'X-Requested-With': 'XMLHttpRequest',\n                            }\n                        };\n                    }\n                    else {\n                        opt.method = method;\n                        opt.headers = __assign({ 'Content-Type': 'application/json', 'Accept': 'application/json', 'X-Requested-With': 'XMLHttpRequest' }, opt.headers);\n                    }\n                    if ('Content-Type' in opt.headers) {\n                        switch (opt.headers['Content-Type']) {\n                            case 'application/json':\n                                opt.body = JSON.stringify(body);\n                                break;\n                            case 'multipart/form-data':\n                                data = new FormData();\n                                for (item in body) {\n                                    data.append(item, body[item]);\n                                }\n                                opt.body = data;\n                                break;\n                        }\n                    }\n                    return [4 /*yield*/, fetch(url, opt)];\n                case 1: return [2 /*return*/, _a.sent()];\n            }\n        });\n    });\n}\n\n\n\n//# sourceURL=webpack://template/./ts/http.ts?");

/***/ }),

/***/ "./ts/subscribe.ts":
/*!*************************!*\
  !*** ./ts/subscribe.ts ***!
  \*************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"subscribeFormHandler\": () => /* binding */ subscribeFormHandler\n/* harmony export */ });\n/* harmony import */ var yup__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! yup */ \"./node_modules/yup/es/index.js\");\n/* harmony import */ var _http__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./http */ \"./ts/http.ts\");\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nvar __generator = (undefined && undefined.__generator) || function (thisArg, body) {\n    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;\n    return g = { next: verb(0), \"throw\": verb(1), \"return\": verb(2) }, typeof Symbol === \"function\" && (g[Symbol.iterator] = function() { return this; }), g;\n    function verb(n) { return function (v) { return step([n, v]); }; }\n    function step(op) {\n        if (f) throw new TypeError(\"Generator is already executing.\");\n        while (_) try {\n            if (f = 1, y && (t = op[0] & 2 ? y[\"return\"] : op[0] ? y[\"throw\"] || ((t = y[\"return\"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;\n            if (y = 0, t) op = [op[0] & 2, t.value];\n            switch (op[0]) {\n                case 0: case 1: t = op; break;\n                case 4: _.label++; return { value: op[1], done: false };\n                case 5: _.label++; y = op[1]; op = [0]; continue;\n                case 7: op = _.ops.pop(); _.trys.pop(); continue;\n                default:\n                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }\n                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }\n                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }\n                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }\n                    if (t[2]) _.ops.pop();\n                    _.trys.pop(); continue;\n            }\n            op = body.call(thisArg, _);\n        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }\n        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };\n    }\n};\n\n\nvar Swal = __webpack_require__(/*! sweetalert2 */ \"./node_modules/sweetalert2/dist/sweetalert2.all.js\");\nvar schema = (0,yup__WEBPACK_IMPORTED_MODULE_0__.object)().shape({\n    name: (0,yup__WEBPACK_IMPORTED_MODULE_0__.string)().required().max(50),\n    email: (0,yup__WEBPACK_IMPORTED_MODULE_0__.string)().required().email().max(150),\n});\nvar subscribe = function (dto) { return __awaiter(void 0, void 0, void 0, function () {\n    var res, data, err_1, validationError_1;\n    return __generator(this, function (_a) {\n        switch (_a.label) {\n            case 0:\n                _a.trys.push([0, 4, , 5]);\n                return [4 /*yield*/, schema.validate(dto, { recursive: true, abortEarly: false })];\n            case 1:\n                _a.sent();\n                return [4 /*yield*/, (0,_http__WEBPACK_IMPORTED_MODULE_1__.http)('/subscribe', 'POST', dto)];\n            case 2:\n                res = _a.sent();\n                return [4 /*yield*/, res.json()];\n            case 3:\n                data = _a.sent();\n                return [2 /*return*/, {\n                        id: data.id,\n                        name: data.name,\n                        email: data.email,\n                        createdAt: new Date(data.createdAt),\n                    }];\n            case 4:\n                err_1 = _a.sent();\n                if (err_1 instanceof yup__WEBPACK_IMPORTED_MODULE_0__.ValidationError) {\n                    validationError_1 = { nameError: '', emailError: '' };\n                    err_1.inner.forEach(function (item) {\n                        if (item.path === 'name') {\n                            validationError_1.nameError = item.errors[0];\n                        }\n                        else if (item.path === 'email') {\n                            validationError_1.emailError = item.errors[0];\n                        }\n                    });\n                    return [2 /*return*/, validationError_1];\n                }\n                return [2 /*return*/, {\n                        message: 'Try again please',\n                    }];\n            case 5: return [2 /*return*/];\n        }\n    });\n}); };\nvar subscribeFormHandler = function (nameEl, emailEl, nameErrorEl, emailErrorEl) { return function (e) { return __awaiter(void 0, void 0, void 0, function () {\n    var name, email, nameError, emailError, res;\n    return __generator(this, function (_a) {\n        switch (_a.label) {\n            case 0:\n                e.preventDefault();\n                name = document.getElementById(nameEl);\n                email = document.getElementById(emailEl);\n                nameError = document.getElementById(nameErrorEl);\n                emailError = document.getElementById(emailErrorEl);\n                return [4 /*yield*/, subscribe({ name: name.value, email: email.value })];\n            case 1:\n                res = _a.sent();\n                // Server error\n                if ('message' in res) {\n                    Swal.fire({\n                        title: 'Error',\n                        text: res.message,\n                        icon: 'error',\n                        timerProgressBar: true,\n                    });\n                    return [2 /*return*/];\n                }\n                // Validation Error\n                if ('nameError' in res && 'emailError' in res) {\n                    if (res.nameError !== '') {\n                        name.classList.add('input-error');\n                        nameError.classList.remove('hidden');\n                        nameError.innerText = res.nameError;\n                    }\n                    if (res.emailError !== '') {\n                        email.classList.add('input-error');\n                        emailError.classList.remove('hidden');\n                        emailError.innerText = res.emailError;\n                    }\n                    setTimeout(function () {\n                        nameError.classList.add('hidden');\n                        emailError.classList.add('hidden');\n                        email.classList.remove('input-error');\n                        name.classList.remove('input-error');\n                    }, 4000);\n                    return [2 /*return*/];\n                }\n                //@ts-ignore\n                gtag('event', 'subscribe', {\n                    event_category: 'subscription',\n                    event_label: 'New user subscribed to news letters',\n                });\n                Swal.fire({\n                    title: 'Success',\n                    text: 'You have successfully subscribed to newsletters',\n                    icon: 'success',\n                    timerProgressBar: true,\n                });\n                return [2 /*return*/];\n        }\n    });\n}); }; };\n\n\n\n//# sourceURL=webpack://template/./ts/subscribe.ts?");

/***/ })

}]);