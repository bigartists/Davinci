/* eslint-disable perfectionist/sort-imports */
import 'src/global.css';

// ----------------------------------------------------------------------

import Router from 'src/routes/sections';

import { useScrollToTop } from 'src/hooks/use-scroll-to-top';

import ThemeProvider from 'src/theme';

import ProgressBar from 'src/components/progress-bar';
import { MotionLazy } from 'src/components/animate/motion-lazy';
import { SettingsDrawer, SettingsProvider } from 'src/components/settings';

import { AuthProvider } from 'src/auth/context/jwt';

// ----------------------------------------------------------------------

export default function App() {
  const charAt = `

  ___              _            _
 |   \\  __ _ __ __(_) _ _   __ (_)
 | |) |/ _\` |\\ V /| || ' \\ / _|| |
 |___/ \\__,_| \\_/ |_||_||_|\\__||_|

  `;

  console.info(`%c${charAt}`, 'color: #5BE49B');

  useScrollToTop();

  return (
    <AuthProvider>
      <SettingsProvider
        defaultSettings={{
          themeMode: 'light', // 'light' | 'dark'
          themeDirection: 'ltr', //  'rtl' | 'ltr'
          themeContrast: 'default', // 'default' | 'bold'
          themeLayout: 'vertical', // 'vertical' | 'horizontal' | 'mini'
          themeColorPresets: 'default', // 'default' | 'cyan' | 'purple' | 'blue' | 'orange' | 'red'
          themeStretch: false,
        }}
      >
        <ThemeProvider>
          <MotionLazy>
            <SettingsDrawer />
            <ProgressBar />
            <Router />
          </MotionLazy>
        </ThemeProvider>
      </SettingsProvider>
    </AuthProvider>
  );
}