import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators';

import { IMusic } from '@/models/music';
import { MusicApi } from '@/api/music';
import store from '..';

@Module
export default class MusicModule extends VuexModule {

    public get musicList(): IMusic[] {
        return this.list;
    }

    public get musicCountVisible(): number {

        let counter = 0;

        for (const music of this.list) {
            if (!music.hidden) {
                counter = counter + 1;
            }
        }

        return counter;
    }

    private static stringIsInclude(field: string, value: string): boolean {

        if (field.toLowerCase().includes(value.toLowerCase())) {
            return true;
        }

        return false;
    }


    public list: IMusic[] = [];

    @Mutation
    public setMusicList(musicList: IMusic[]) {
        this.list = musicList;
    }

    @Action({ commit: 'setMusicList' })
    public async fetchAll() {
        const musicApi = MusicApi.GetInstance();
        const response = await musicApi.GetAllMusic();
        if (musicApi.UpdatingErrMsg !== '') {
            return this.list;
        }

        return typeof response === 'undefined' ? this.list : response;
    }

    @Action({})
    public async addMusic(data: IMusic) {
        const musicApi = MusicApi.GetInstance();
        await musicApi.InsertNewMusic(data);
        store.dispatch('fetchAll');
    }

    @Action({ commit: "setMusicList" })
    public filterResult(searchString: string) {

        const result: IMusic[] = [];

        const musicList: IMusic[] = JSON.parse(JSON.stringify(this.list));

        for (const music of musicList) {

            music.hidden = false;

            // tslint:disable-next-line:max-line-length
            if (!(MusicModule.stringIsInclude(music.artist, searchString) || MusicModule.stringIsInclude(music.title, searchString))) {
                music.hidden = true;
            }

            result.push(music);
        }

        return result;
    }

}

