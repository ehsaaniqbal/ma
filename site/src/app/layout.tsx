import "./globals.css";
import { Inter } from "next/font/google";
import Script from "next/script";
import { Analytics } from "@vercel/analytics/react";
import { isMobile } from "react-device-detect";

const inter = Inter({ subsets: ["latin"], weight: [] });

export const metadata = {
  title: "",
  description: "",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  if (isMobile) {
    alert(
      "This site isn't optimized for mobile view. Probably never will be lol"
    );
  }

  return (
    <html lang="en">
      <body className={inter.className}>
        {children}
        <Script src="/wasm_exec.js" strategy="beforeInteractive" />
        <Analytics />
      </body>
    </html>
  );
}
