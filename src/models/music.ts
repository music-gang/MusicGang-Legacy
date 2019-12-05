export interface IMusic {
    title: string;
    artist: string;
    genre?: string;
    notes?: string;
    prompter?: string;
}

export class MusicDTO implements IMusic {
    title: string = "";
    artist: string = "";
    genre?: string;
    notes?: string;
    prompter?: string;
}

export class Music extends MusicDTO {

    constructor(dto: MusicDTO) {
        super()
        Object.assign(this, dto)
    }
}