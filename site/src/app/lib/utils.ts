export const parseMaOutput = (code: string) => {
  const parsed = JSON.parse(code);
  return parsed.map((op: Record<string, string>) => {
    console.log(op);
    const key = Object.keys(op)[0];
    switch (key) {
      case "string":
        return `<p style="font-weight: 500"> ${op[key]}</p>`;
      case "error":
        return `<p style="color: red; font-weight: 500"> ${op[key]}</p>`;
    }
  });
};
