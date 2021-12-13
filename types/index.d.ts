interface TarHeader {
  Name: string;
  Size: string;
  Linkname: string;
  Mode: number;
  Uid: number;
  Gid: number;
  Uname: string;
  Gname: string;
  ModTime: Date;
  AccessTime: Date;
  ChangeTime: Date;
  Format: number;
}

type Callback = (header: TarHeader) => boolean;

export declare function targz(a: Blob, callback: Callback): number;
