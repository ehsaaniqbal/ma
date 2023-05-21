import React from "react";
import { highlight } from "prismjs/components/prism-core";
import "prismjs/themes/prism-tomorrow.css";
import { maSyntax } from "../lib/syntax";

const docContent = [
  {
    title: "Types",
    description:
      "Discover the seamless harmony of null, bool, int, str, and the audacious initiatives (functions). Embrace the limitless possibilities as you shape the future of programming with MaScript's curated collection.",
    code: `ma LPA = 3
ma category = "super dream"
ma promotion = null
ma happiness = false
`,
  },
  {
    title: "Arithmetic",
    description:
      "Experience a seamless blend of style and precision as addition, subtraction, multiplication, and division redefine the art of calculations. Discover the beauty and ingenuity that lies within each operation, brought to you by the visionary minds at MaScript.",
    code: `ma skills = 404
ma marks = 16
ma future = skills + marks
// future = 420
`,
  },
  {
    title: "Conditional Expressions",
    description:
      "Embrace the power of 'if' and 'else', navigating possibilities with elegance. Let MaScript redefine decision-making, illuminating the path to coding mastery.",
    code: `super dream (immersion == true)  {
    roar("lion")
} regular {
    roar("donkey")
}
`,
  },
  {
    title: "Initiatives",
    description:
      "MaScript's symphony of creativity. Unleash the power of named or anonymous functions, including nested closures. Embrace the artistry as you compose elegant and harmonious code with MaScript's visionary design.",
    code: `ma advitya = initiative (money) {
    ma fun = false
    ma creativity = null
    god bless you money / 0
}
`,
  },
  {
    title: "Roar",
    description:
      "Feel the thunderous impact of 'roar' and embrace the orchestration of MaScript's built-in function ensemble.",
    code: `roar("Work hard - success will be yours.")`,
  },
];

const Doc = ({
  title,
  description,
  code,
  style,
}: {
  title: string;
  description: string;
  code: string;
  style: string;
}) => {
  return (
    <div
      className={`flex flex-col justify-between rounded-lg border border-transparent px-5 py-4 border-neutral-700 hover:bg-neutral-800/30 ${style}`}
    >
      <h2 className="leading-[2] tracking-tighter font-extrabold text-3xl bg-[linear-gradient(180deg,_#fff,_#adadad)] bg-clip-text text-transparent">
        {title}
      </h2>
      <p className={`my-3 text-md opacity-60 text-justify max-w-md text-md`}>
        {description}
      </p>
      <pre>
        <code
          className="language-ma"
          dangerouslySetInnerHTML={{ __html: highlight(code, maSyntax, "ma") }}
        />
      </pre>
    </div>
  );
};

const Docs = () => {
  return (
    <section
      id="docs"
      className="flex min-h-screen flex-col items-center px-24 mb-5"
    >
      <div className="p-4">
        <h1 className="leading-[2] tracking-tighter font-extrabold text-6xl bg-[linear-gradient(180deg,_#fff,_#adadad)] bg-clip-text text-transparent">
          Marauders Map to MaScript
        </h1>
      </div>
      <div className="grid text-center lg:mb-0 lg:grid-cols-2 lg:text-left gap-5">
        {docContent.map((dc, index) => (
          <Doc
            code={dc.code}
            title={dc.title}
            description={dc.description}
            style={
              index === docContent.length - 1
                ? "col-span-2 w-[60%] justify-self-center"
                : ""
            }
          />
        ))}
      </div>
    </section>
  );
};

export default Docs;
