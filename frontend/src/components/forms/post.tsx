import {useState} from "react";
import OutputPost from "./output-post";

function Post() {
    const [text, setText] = useState('')
    const [header, setHeader] = useState('')
    const textInput = <textarea rows={8} cols={55} name="text" id="form_textarea" onChange={(v)=> {
        setText(v.target.value)
    }} />
    return (
        <>
            <form id="f1" action="functions/new_post.php" method="post" >
                <div className="form">
                    <span className="form_name" id="form_name"></span>
                    <div className="form_smiles" id="smiles_1">
                        <div id="banner"></div>
                    </div>
                    <input className="form_title" type="text" name="title" placeholder="Заголовок" onChange={(v)=> setHeader(v.target.value)}/>
                    <input type="text" name="email" />
                    <span className="form_BB_codes" onClick={() => console.log("BB")
                    } title="Полужирный">
                        <strong>B</strong>
                    </span>
                    <span className="form_BB_codes" title="Курсив">
                        <em>I</em>
                    </span>
                    <span className="form_BB_codes" title="Зачеркнутый">
                        <del>S</del>
                    </span>
                    <span className="form_BB_codes" title="Подчеркнутый">
                        <ins>U</ins>
                    </span>
                    <span className="form_BB_codes" title="Спойлер">Sp</span>
                    <span className="form_BB_codes" title="Цитата">Q</span>
                    <span className="form_BB_codes" title="Смайлы">:)</span>
                    {textInput}
                    <input type="file" className="form_pc_file" name="img" />
                    <input type="hidden" id="replay_id" name="id" value="" />
                    <input type="submit" className="form_submit" value="Отправить" />
                </div>
            </form>
            <OutputPost text={text} header={header}></OutputPost>
            </>
    );
}

export default Post;