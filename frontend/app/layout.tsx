import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { Navbar } from '@/components/layout/navbar';
import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Event Management App',
  description: 'Manage and discover events',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="min-h-screen flex flex-col">
          <Navbar />
          <main className="flex-1">{children}</main>
          <footer className="border-t py-6 mt-12">
            <div className="container mx-auto px-4 text-center text-sm text-muted-foreground">
              Event Management App &copy; {new Date().getFullYear()}
            </div>
          </footer>
        </div>
      </body>
    </html>
  );
}
