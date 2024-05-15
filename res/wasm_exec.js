(() => {
    if ("undefined" != typeof global) ; else if ("undefined" != typeof window) window.global = window; else if ("undefined" != typeof self) self.global = self; else throw Error("cannot export Go (neither global, window nor self is defined)");
    global.require || "undefined" == typeof require || (global.require = require), !global.fs && global.require && (global.fs = require("fs"));
    let e = () => {
        let e = Error("not implemented");
        return e.code = "ENOSYS", e
    };
    if (!global.fs) {
        let t = "";
        global.fs = {
            constants: {O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1},
            writeSync(e, s) {
                t += r.decode(s);
                let n = t.lastIndexOf("\n");
                return -1 != n && (console.log(t.substr(0, n)), t = t.substr(n + 1)), s.length
            },
            write(t, s, n, r, i, l) {
                if (0 !== n || r !== s.length || null !== i) {
                    l(e());
                    return
                }
                let o = this.writeSync(t, s);
                l(null, o)
            },
            chmod(t, s, n) {
                n(e())
            },
            chown(t, s, n, r) {
                r(e())
            },
            close(t, s) {
                s(e())
            },
            fchmod(t, s, n) {
                n(e())
            },
            fchown(t, s, n, r) {
                r(e())
            },
            fstat(t, s) {
                s(e())
            },
            fsync(e, t) {
                t(null)
            },
            ftruncate(t, s, n) {
                n(e())
            },
            lchown(t, s, n, r) {
                r(e())
            },
            link(t, s, n) {
                n(e())
            },
            lstat(t, s) {
                s(e())
            },
            mkdir(t, s, n) {
                n(e())
            },
            open(t, s, n, r) {
                r(e())
            },
            read(t, s, n, r, i, l) {
                l(e())
            },
            readdir(t, s) {
                s(e())
            },
            readlink(t, s) {
                s(e())
            },
            rename(t, s, n) {
                n(e())
            },
            rmdir(t, s) {
                s(e())
            },
            stat(t, s) {
                s(e())
            },
            symlink(t, s, n) {
                n(e())
            },
            truncate(t, s, n) {
                n(e())
            },
            unlink(t, s) {
                s(e())
            },
            utimes(t, s, n, r) {
                r(e())
            }
        }
    }
    if (global.process || (global.process = {
        getuid: () => -1,
        getgid: () => -1,
        geteuid: () => -1,
        getegid: () => -1,
        getgroups() {
            throw e()
        },
        pid: -1,
        ppid: -1,
        umask() {
            throw e()
        },
        cwd() {
            throw e()
        },
        chdir() {
            throw e()
        }
    }), !global.crypto) {
        let s = require("crypto");
        global.crypto = {
            getRandomValues(e) {
                s.randomFillSync(e)
            }
        }
    }
    global.performance || (global.performance = {
        now() {
            let [e, t] = process.hrtime();
            return 1e3 * e + t / 1e6
        }
    }), global.TextEncoder || (global.TextEncoder = require("util").TextEncoder), global.TextDecoder || (global.TextDecoder = require("util").TextDecoder);
    let n = new TextEncoder("utf-8"), r = new TextDecoder("utf-8"), i = new DataView(new ArrayBuffer(8));
    var l = [];
    if (global.Go = class {
        constructor() {
            this._callbackTimeouts = new Map, this._nextCallbackTimeoutID = 1;
            let e = () => new DataView(this._inst.exports.memory.buffer), t = e => {
                    i.setBigInt64(0, e, !0);
                    let t = i.getFloat64(0, !0);
                    if (0 === t) return;
                    if (!isNaN(t)) return t;
                    let s = 4294967295n & e;
                    return this._values[s]
                }, s = s => {
                    let n = e().getBigUint64(s, !0);
                    return t(n)
                }, o = e => {
                    if ("number" == typeof e) return isNaN(e) ? 2146959360n << 32n : 0 === e ? 2146959360n << 32n | 1n : (i.setFloat64(0, e, !0), i.getBigInt64(0, !0));
                    switch (e) {
                        case void 0:
                            return 0n;
                        case null:
                            return 2146959360n << 32n | 2n;
                        case!0:
                            return 2146959360n << 32n | 3n;
                        case!1:
                            return 2146959360n << 32n | 4n
                    }
                    let t = this._ids.get(e);
                    void 0 === t && (void 0 === (t = this._idPool.pop()) && (t = BigInt(this._values.length)), this._values[t] = e, this._goRefCounts[t] = 0, this._ids.set(e, t)), this._goRefCounts[t]++;
                    let s = 1n;
                    switch (typeof e) {
                        case"string":
                            s = 2n;
                            break;
                        case"symbol":
                            s = 3n;
                            break;
                        case"function":
                            s = 4n
                    }
                    return t | (2146959360n | s) << 32n
                }, a = (t, s) => {
                    let n = o(s);
                    e().setBigUint64(t, n, !0)
                }, c = (e, t, s) => new Uint8Array(this._inst.exports.memory.buffer, e, t), u = (e, t, n) => {
                    let r = Array(t);
                    for (let i = 0; i < t; i++) r[i] = s(e + 8 * i);
                    return r
                }, $ = (e, t) => r.decode(new DataView(this._inst.exports.memory.buffer, e, t)),
                f = Date.now() - performance.now();
            this.importObject = {
                wasi_snapshot_preview1: {
                    fd_write: function (t, s, n, i) {
                        let o = 0;
                        if (1 == t) for (let a = 0; a < n; a++) {
                            let c = s + 8 * a, u = e().getUint32(c + 0, !0), $ = e().getUint32(c + 4, !0);
                            o += $;
                            for (let f = 0; f < $; f++) {
                                let d = e().getUint8(u + f);
                                if (13 == d) ; else if (10 == d) {
                                    let h = r.decode(new Uint8Array(l));
                                    l = [], console.log(h)
                                } else l.push(d)
                            }
                        } else console.error("invalid file descriptor:", t);
                        return e().setUint32(i, o, !0), 0
                    }, fd_close: () => 0, fd_fdstat_get: () => 0, fd_seek: () => 0, proc_exit(e) {
                        if (global.process) process.exit(e); else throw "trying to exit with code " + e
                    }, random_get: (e, t) => (crypto.getRandomValues(c(e, t)), 0)
                }, gojs: {
                    "runtime.ticks": () => f + performance.now(),
                    "runtime.sleepTicks": e => {
                        setTimeout(this._inst.exports.go_scheduler, e)
                    },
                    "syscall/js.finalizeRef": (v_ref) => {
                        // TODO: Fix
                        // const id = mem().getUint32(unboxValue(v_ref), true);
                        // this._goRefCounts[id]--;
                        // if (this._goRefCounts[id] === 0) {
                        //     const v = this._values[id];
                        //     this._values[id] = null;
                        //     this._ids.delete(v);
                        //     this._idPool.push(id);
                        // }
                    },
                    "syscall/js.stringVal"(e, t) {
                        let s = $(e, t);
                        return o(s)
                    },
                    "syscall/js.valueGet"(e, s, n) {
                        let r = $(s, n), i = t(e), l = Reflect.get(i, r);
                        return o(l)
                    },
                    "syscall/js.valueSet"(e, s, n, r) {
                        let i = t(e), l = $(s, n), o = t(r);
                        Reflect.set(i, l, o)
                    },
                    "syscall/js.valueDelete"(e, s, n) {
                        let r = t(e), i = $(s, n);
                        Reflect.deleteProperty(r, i)
                    },
                    "syscall/js.valueIndex": (e, s) => o(Reflect.get(t(e), s)),
                    "syscall/js.valueSetIndex"(e, s, n) {
                        Reflect.set(t(e), s, t(n))
                    },
                    "syscall/js.valueCall"(s, n, r, i, l, o, c) {
                        let f = t(n), d = $(r, i), h = u(l, o, c);
                        try {
                            let g = Reflect.get(f, d);
                            a(s, Reflect.apply(g, f, h)), e().setUint8(s + 8, 1)
                        } catch (p) {
                            a(s, p), e().setUint8(s + 8, 0)
                        }
                    },
                    "syscall/js.valueInvoke"(s, n, r, i, l) {
                        try {
                            let o = t(n), c = u(r, i, l);
                            a(s, Reflect.apply(o, void 0, c)), e().setUint8(s + 8, 1)
                        } catch ($) {
                            a(s, $), e().setUint8(s + 8, 0)
                        }
                    },
                    "syscall/js.valueNew"(s, n, r, i, l) {
                        let o = t(n), c = u(r, i, l);
                        try {
                            a(s, Reflect.construct(o, c)), e().setUint8(s + 8, 1)
                        } catch ($) {
                            a(s, $), e().setUint8(s + 8, 0)
                        }
                    },
                    "syscall/js.valueLength": e => t(e).length,
                    "syscall/js.valuePrepareString"(s, r) {
                        let i = String(t(r)), l = n.encode(i);
                        a(s, l), e().setInt32(s + 8, l.length, !0)
                    },
                    "syscall/js.valueLoadString"(e, s, n, r) {
                        let i = t(e);
                        c(s, n, r).set(i)
                    },
                    "syscall/js.valueInstanceOf": (e, s) => t(e) instanceof t(s),
                    "syscall/js.copyBytesToGo"(s, n, r, i, l) {
                        let o = s + 4, a = c(n, r), u = t(l);
                        if (!(u instanceof Uint8Array || u instanceof Uint8ClampedArray)) {
                            e().setUint8(o, 0);
                            return
                        }
                        let $ = u.subarray(0, a.length);
                        a.set($), e().setUint32(s, $.length, !0), e().setUint8(o, 1)
                    },
                    "syscall/js.copyBytesToJS"(s, n, r, i, l) {
                        let o = s + 4, a = t(n), u = c(r, i);
                        if (!(a instanceof Uint8Array || a instanceof Uint8ClampedArray)) {
                            e().setUint8(o, 0);
                            return
                        }
                        let $ = u.subarray(0, a.length);
                        a.set($), e().setUint32(s, $.length, !0), e().setUint8(o, 1)
                    }
                }
            }, this.importObject.env = this.importObject.gojs
        }

        async run(e) {
            for (this._inst = e, this._values = [NaN, 0, null, !0, !1, global, this,], this._goRefCounts = [], this._ids = new Map, this._idPool = [], this.exited = !1; ;) {
                let t = new Promise(e => {
                    this._resolveCallbackPromise = () => {
                        if (this.exited) throw Error("bad callback: Go program has already exited");
                        setTimeout(e, 0)
                    }
                });
                if (this._inst.exports._start(), this.exited) break;
                await t
            }
        }

        _resume() {
            if (this.exited) throw Error("Go program has already exited");
            this._inst.exports.resume(), this.exited && this._resolveExitPromise()
        }

        _makeFuncWrapper(e) {
            let t = this;
            return function () {
                let s = {id: e, this: this, args: arguments};
                return t._pendingEvent = s, t._resume(), s.result
            }
        }
    }, global.require && global.require.main === module && global.process && global.process.versions && !global.process.versions.electron) {
        3 != process.argv.length && (console.error("usage: go_js_wasm_exec [wasm binary] [arguments]"), process.exit(1));
        let o = new Go;
        WebAssembly.instantiate(fs.readFileSync(process.argv[2]), o.importObject).then(e => o.run(e.instance)).catch(e => {
            console.error(e), process.exit(1)
        })
    }
})();