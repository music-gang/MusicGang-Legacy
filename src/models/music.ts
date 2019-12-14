export interface IMusic {
    id?: number;
    title: string;
    artist: string;
    genre?: string;
    notes?: string;
    prompter?: string;
    hidden?: boolean;
}

export class MusicDTO implements IMusic {
    public id?: number;
    public title: string = '';
    public artist: string = '';
    public genre?: string;
    public notes?: string;
    public prompter?: string;
    public hidden?: boolean = false;
}

// tslint:disable-next-line:max-classes-per-file
export class Music extends MusicDTO {

    constructor(dto: MusicDTO) {
        super();
        Object.assign(this, dto);
    }
}
