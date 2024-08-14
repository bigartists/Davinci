import type { IUser } from '../user/type';

export interface IProject {
  createBy: IUser;
  description: string;
  id: number;
  initialOrgId: number;
  isStar: boolean;
  isTransfer: boolean;
  name: string;
  orgId: number;
  permission: IPermission;
  pic: string;
  starNum: number;
  userId: number;
  visibility: boolean;
}

export interface IPermission {
  downloadPermission: boolean;
  sharePermission: boolean;
  schedulePermission: number;
  sourcePermission: number;
  viewPermission: number;
  vizPermission: number;
  widgetPermission: number;
}
