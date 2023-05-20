export default function Home() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center p-24">
        <div className="p-4">
          <h1 className="leading-[2] tracking-tighter font-extrabold text-6xl bg-[linear-gradient(180deg,_#fff,_#adadad)] bg-clip-text text-transparent">
            The Future of Programming is here.
          </h1>
          <p className="text-center">Powered by AI</p>
        </div>
        <div className="mb-32 grid text-center lg:mb-0 lg:grid-cols-2 lg:text-left mt-20">
          <a
            href="#"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            target="_blank"
            rel="noopener noreferrer"
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
            href="#"
            className="group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
            target="_blank"
            rel="noopener noreferrer"
          >
            <h2 className={`mb-3 text-2xl font-semibold`}>
              Docs{" "}
              <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
                -&gt;
              </span>
            </h2>
            <p className={`m-0 max-w-[30ch] text-sm opacity-50`}>
              Learn more about MaScript's awesome features
            </p>
          </a>
        </div>
      </main>
    </>
  );
}
