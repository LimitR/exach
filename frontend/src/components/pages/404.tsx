import logo from '../../img/logo.png';

export default function NotFound(props: {errorText: string}) {
    return (
        <div className="comment">
            <h2>Страница не найдена</h2>
            <img src={logo} alt="" className="text_image" />
            <p className="text">
                {props.errorText ?? "Error"}
            </p>
            <span className="info">&nbsp;</span>
        </div>
    );
}