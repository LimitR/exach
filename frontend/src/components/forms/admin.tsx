
import logo from '../../img/logo.png';
export default function Admin() {
    return (
        <form action="" method="post" encType="multipart/form-data">
            <div className="post">
                <h2>Вход в админку</h2>
                <img src={logo} alt="" className="text_image" />
                <input type="text" name="admin_name" maxLength={15} placeholder="Имя" /><br /><br />
                <input type="password" name="admin_password" maxLength={15} placeholder="Пароль" /><br/>
                <input type="submit" className="form_submit" value="Ok" />
                <div className="info">&nbsp;</div>
            </div>
        </form>
    )
}