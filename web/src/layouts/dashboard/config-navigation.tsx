import { useMemo } from 'react';

// import { paths } from 'src/routes/paths';

import SvgColor from 'src/components/svg-color';

// ----------------------------------------------------------------------

const icon = (name: string) => (
  <SvgColor src={`/assets/icons/navbar/${name}.svg`} sx={{ width: 1, height: 1 }} />
  // OR
  // <Iconify icon="fluent:mail-24-filled" />
  // https://icon-sets.iconify.design/solar/
  // https://www.streamlinehq.com/icons
);

const ICONS = {
  job: icon('ic_job'),
  blog: icon('ic_blog'),
  chat: icon('ic_chat'),
  mail: icon('ic_mail'),
  user: icon('ic_user'),
  file: icon('ic_file'),
  lock: icon('ic_lock'),
  tour: icon('ic_tour'),
  order: icon('ic_order'),
  label: icon('ic_label'),
  blank: icon('ic_blank'),
  kanban: icon('ic_kanban'),
  folder: icon('ic_folder'),
  banking: icon('ic_banking'),
  booking: icon('ic_booking'),
  invoice: icon('ic_invoice'),
  product: icon('ic_product'),
  calendar: icon('ic_calendar'),
  disabled: icon('ic_disabled'),
  external: icon('ic_external'),
  menuItem: icon('ic_menu_item'),
  ecommerce: icon('ic_ecommerce'),
  analytics: icon('ic_analytics'),
  dashboard: icon('ic_dashboard'),
};

// ----------------------------------------------------------------------

export function useNavData() {
  const data = useMemo(
    () => [
      // OVERVIEW
      // ----------------------------------------------------------------------
      // {
      //   subheader: 'overview v5.7.0',
      //   items: [
      //     { title: 'one', path: paths.dashboard.root, icon: ICONS.dashboard },
      //     { title: 'two', path: paths.dashboard.two, icon: ICONS.ecommerce },
      //     {
      //       title: 'three',
      //       path: paths.dashboard.three,
      //       icon: ICONS.analytics,
      //     },
      //   ],
      // },

      // MANAGEMENT
      // ----------------------------------------------------------------------
      // {
      //   subheader: 'management',
      //   items: [
      //     {
      //       title: 'user',
      //       path: paths.dashboard.group.root,
      //       icon: ICONS.user,
      //       children: [
      //         { title: 'four', path: paths.dashboard.group.root },
      //         { title: 'five', path: paths.dashboard.group.five },
      //         { title: 'six', path: paths.dashboard.group.six },
      //       ],
      //     },
      //   ],
      // },

      {
        subheader: 'Workload',
        items: [
          { title: 'Dashboard', path: '/', icon: ICONS.dashboard },
          // { title: 'Deployment', path: '/deployment', icon: ICONS.folder },
          // { title: 'Pod', path: '/pod', icon: ICONS.ecommerce },
          // { title: 'Service', path: '/service', icon: ICONS.analytics },
        ],
      },

      // {
      //   subheader: 'Storage',
      //   items: [
      //     { title: 'ConfigMap', path: '/configmap', icon: ICONS.folder },
      //     { title: 'Secret', path: '/secret', icon: ICONS.kanban },
      //   ],
      // },

      // {
      //   subheader: 'Application',
      //   items: [
      //     {
      //       title: 'Application',
      //       path: '/app',
      //       icon: ICONS.user,
      //       children: [
      //         { title: 'Login', path: '/app/one' },
      //         { title: 'Not Found', path: '/app/two' },
      //       ],
      //     },
      //   ],
      // },
    ],
    []
  );

  return data;
}
