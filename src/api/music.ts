import Axios from 'axios';
import { IMusic, Music } from '@/models/music';

interface MusicRequestInterface {
    status: number;
    success: boolean;
    message: string;
    content?: IMusic[];
}

export class MusicApi {

    //#region static

    public static GetInstance(): MusicApi {

        if (typeof MusicApi.instance === 'undefined') {
            MusicApi.instance = new MusicApi();
        }

        return MusicApi.instance;
    }

    private static instance?: MusicApi = undefined;

    private static musicAxios = Axios.create();

    //#endregion

    //#region instance

    private updating: boolean = false;

    private updatingErrMsg: string = '';

    get Updating(): boolean {
        return this.Updating;
    }

    get UpdatingErrMsg(): string {
        return this.updatingErrMsg;
    }

    public async GetAllMusic(): Promise<Music[] | void> {

        if (this.updating) {
            return;
        }

        this.updating = true;
        this.updatingErrMsg = '';

        const url = '/music/all';
        const response = await MusicApi.musicAxios.get<MusicRequestInterface>(url);

        this.updating = false;

        const musicList = response.data.content;
        if (response.data.status !== 0 || !response.data.success || typeof musicList === 'undefined') {
            this.updatingErrMsg = 'Errore durante il recupero della lista delle canzoni';
            return [];
        }

        return musicList.map((musicDto) => new Music(musicDto));
    }


    public async InsertNewMusic(music: IMusic): Promise<boolean | void> {

        if (this.updating) {
            return;
        }

        this.updating = true;
        this.updatingErrMsg = '';

        const url = 'music/new';
        const response = await MusicApi.musicAxios.post<MusicRequestInterface>(url, music);

        if (response.data.status !== 0 || !response.data.success) {
            this.updatingErrMsg = 'Errore durante il recupero della lista delle canzoni';
            return false;
        }

        return true;
    }

    //#endregion
}
