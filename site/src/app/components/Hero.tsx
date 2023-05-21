import React from "react";

const Hero = () => {
  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <div className="p-4">
        <h1 className="leading-[2] tracking-tighter font-extrabold text-7xl bg-[linear-gradient(180deg,_#fff,_#adadad)] bg-clip-text text-transparent">
          The Future of Programming is here.
        </h1>
      </div>
      <p className="text-center text-[#888] text-xl p-4">
        The Quantum-Powered Blockchain AI Language that Unleashes God like
        Powers! Say Goodbye to the Langauges of the Past and Code the Future
        with <span className="text-white font-semibold">MaScript</span>
      </p>

      <div className="mb-32 grid text-center lg:mb-0 lg:grid-cols-2 lg:text-left mt-20">
        <a
          href="#playground"
          className="group rounded-lg border border-transparent px-5 py-4 transition-colors border-neutral-700 hover:bg-neutral-800/30"
        >
          <h2 className={`mb-3 text-2xl font-semibold`}>
            Playground{" "}
            <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
              -&gt;
            </span>
          </h2>
          <p className={`m-0 max-w-[30ch] text-sm opacity-50`}>
            Get your hands dirty with MaScript
          </p>
        </a>
        <a
          href="#docs"
          className="group rounded-lg border border-transparent px-5 py-4 transition-colors border-neutral-700 hover:bg-neutral-800/30"
        >
          <h2 className={`mb-3 text-2xl font-semibold`}>
            Docs{" "}
            <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
              -&gt;
            </span>
          </h2>
          <p className={`m-0 max-w-[30ch] text-sm opacity-50`}>
            Learn more about MaScript&apos;s revolutionary features
          </p>
        </a>
      </div>
    </main>
  );
};

export default Hero;
