import ReactMarkdown from "react-markdown";



export default function OutputPost(
    props: {
        header: string,
        text: string,
    }
) {
    return (
        <div className="post" id="c1">
            <ReactMarkdown>{"# " + props.header}</ReactMarkdown> <p><ReactMarkdown>{props.text}</ReactMarkdown></p>
        </div>
    )
}