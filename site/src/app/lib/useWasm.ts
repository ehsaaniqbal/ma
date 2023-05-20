import { useState, useEffect } from "react";

export function useWasm(path: string) {
  const [state, setState] = useState<[any, boolean, Error | null]>([
    null,
    true,
    null,
  ]);

  async function getWasm(path: string) {
    try {
      // @ts-ignore
      const wasm_exec = window?.Go;
      if (wasm_exec) {
        // @ts-ignore
        const go = new Go(); // Defined in wasm_exec.js
        const WASM_URL = path;

        let wasm;

        if ("instantiateStreaming" in WebAssembly) {
          const obj = await WebAssembly.instantiateStreaming(
            fetch(WASM_URL),
            go.importObject
          );

          wasm = obj.instance;
          go.run(wasm);

          return wasm.exports;
        } else {
          const resp = await fetch(WASM_URL);
          const bytes = await resp.arrayBuffer();
          const obj = await WebAssembly.instantiate(bytes, go.importObject);

          wasm = obj.instance;
          go.run(wasm);

          return wasm.exports;
        }
      }
    } catch (e) {
      console.log(e);
      return {};
    }
  }

  useEffect(() => {
    getWasm(path)
      .then((exp) => {
        setState([exp, false, null]);
        console.log("wasm initialized");
      })
      .catch((err) => {
        console.error(err);
        setState([null, false, err]);
      });
  }, [path]);

  return state;
}
