import React from "react";
import { parseMaOutput } from "../lib/utils";

interface TerminalProps {
  code: string;
}

const Terminal = ({ code }: TerminalProps) => {
  return (
    <div
      id="terminal"
      className="rounded-lg border-solid border-[1px] border-[#2e2e2e] w-[60vw] h-[30vh] mt-2 p-4 overflow-scroll"
    >
      <div className="border-solid border-b-[1px]">
        <h4 className="text-2xl font-semibold">Output</h4>
      </div>
      <div className="mt-4">
        {parseMaOutput(code).map((op: string) => {
          console.log(op);
          return <div dangerouslySetInnerHTML={{ __html: op }} />;
        })}
      </div>
    </div>
  );
};

export default Terminal;
