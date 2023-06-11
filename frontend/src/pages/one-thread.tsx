import { useState, useEffect } from "react";
import {useFetching} from '../hooks/useFetching'
import axios from "axios";

type Thread = {
    id: number,
    text: string,
    head: string,
    img: string
}


export default function OneThread() {
    const [data, setData]: [ Thread[], any ] = useState([{}] as Thread[]);

    const [fetching, isLoad, _] = useFetching(async ()=>{
        const res = await axios.get<Thread[]>('http://localhost/api/thread/10')
        setData(res.data)
    })
    useEffect(() => {
        //@ts-ignore
        fetching()
    }, [])

    if (!data.length) {
        return (
            <div>
                NOT TABS
            </div>
        )
    }

    return (
        <div>
            {isLoad ? "Загрузка" : null}
            {data.map(e => (
                <div>
                    <div>{e.id}</div>
                    <div>{e.head}</div>
                    <div>{e.text}</div>
                    <div>{e.img}</div>
                </div>
            ))}
        </div>
    );
}