import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Echo + EdgeOne Pages",
  description: "Go Functions allow you to run Go web frameworks like Echo on EdgeOne Pages. Build full-stack applications with Echo's high-performance routing and middleware.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en-US">
      <head>
        <link rel="icon" href="/echo-favicon.svg" />
      </head>
      <body
        className="antialiased"
      >
        {children}
      </body>
    </html>
  );
}
