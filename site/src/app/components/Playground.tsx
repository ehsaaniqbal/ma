"use client";

import "prismjs/themes/prism-tomorrow.css";
import React, { useState, useRef, useEffect } from "react";
import Editor from "react-simple-code-editor";
import { highlight } from "prismjs";
import { maSyntax } from "../lib/syntax";
import Terminal from "./Terminal";
import Button from "./Button";
import { useWasm } from "../lib/useWasm";
import { isMobile } from "react-device-detect";

const maBasic = `ma x = 34
ma y = 2

ma mul = initiative(x,y) {
    x * y
}

roar(mul(x,y) + 1)

super dream (x > y) {
    roar("NO PLACEMENT FOR YOU MA!")
} regular {
    roar("STILL NO PLACEMENT")
}
`;

const Playground = () => {
  useWasm("/main.wasm");

  const [code, setCode] = useState(maBasic);
  const [output, setOutput] = useState("");
  const [showTerminal, setShowTerminal] = useState(false);

  let maEval = useRef(null);
  let maOutput = useRef(null);

  if (process.browser) {
    //@ts-ignore
    maEval.current = window?.ma_eval;
    //@ts-ignore
    maOutput.current = window?.ma_output;
  }

  useEffect(() => {
    if (isMobile) {
      alert(
        "This site isn't optimized for mobile view and never will be lol. Use a desktop for the best experience"
      );
    }
  }, []);

  return (
    <section
      className="flex min-h-screen flex-col items-center px-24"
      id="playground"
    >
      <div className="p-4">
        <h1 className="leading-[2] tracking-tighter font-extrabold text-6xl bg-[linear-gradient(180deg,_#fff,_#adadad)] bg-clip-text text-transparent">
          Try MaScript Out
        </h1>
      </div>
      <div className="whitespace-nowrap overflow-x-hidden flex justify-center items-center language-ma color-[#ccc]">
        <Editor
          value={code}
          onValueChange={(code) => setCode(code)}
          highlight={(code) => highlight(code, maSyntax, "ma")}
          padding={10}
          style={{
            fontFamily:
              "Consolas, Monaco, 'Andale Mono', 'Ubuntu Mono', monospace",
            overflowY: "scroll",
            overflowX: "hidden",
            wordBreak: "keep-all",
            height: "70vh",
            width: "60vw",
          }}
          className="line-numbers flex-none whitespace-nowrap editor overflow-y-scroll overflow-x-scroll rounded-md text-[#ccc] bg-[#2d2d2d]"
          textareaClassName="overflow-x-hidden whitespace-nowrap"
          textareaId="codeArea"
        />
      </div>
      <div className="my-4 flex flex-row w-[40vw] justify-center">
        <Button
          filled={true}
          label="Run"
          href="#terminal"
          onClick={() => {
            setShowTerminal(true);
            if (maEval?.current) {
              // @ts-ignore
              maEval.current(`${code}`);
              // @ts-ignore
              setOutput(window?.ma_output);
            }
          }}
        />
        <Button
          filled={false}
          label="Clear"
          href="#playground"
          onClick={() => {
            setShowTerminal(false);
          }}
        />
      </div>
      {showTerminal ? <Terminal code={output} /> : null}
    </section>
  );
};

export default Playground;
