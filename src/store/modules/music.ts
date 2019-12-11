import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators';

import { IMusic } from '@/models/music';
import { MusicApi } from '@/api/music';
import store from '..';

@Module
export default class MusicModule extends VuexModule {

    public list: IMusic[] = [];

    public get musicList(): IMusic[] {
        return this.list;
    }

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

        return response;
    }
    @Action({})
    public async addMusic(data: IMusic) {
        const musicApi = MusicApi.GetInstance();
        await musicApi.InsertNewMusic(data);
        store.dispatch('fetchAll');
    }
}
