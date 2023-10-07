import '../styles/globals.css';
import Nav from './components/Nav/Navbar';
import { Providers } from './providers';

export default function RootLayout({ children }) {
  return (
    <html lang="en" className="dark">
      <body>
        <Providers>
          <Nav />
          {children}
        </Providers>
      </body>
    </html>
  );
}
