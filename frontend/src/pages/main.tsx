
import Menu from '../components/menu';
import logo from '../img/logo.png';
import Footer from './footer';

export default function Main(props?: { page: any }) {
    const content = () => {
        if (!props?.page) {
            return (
                <div className="post">
                    <img src={logo} className="text_image" alt="\" />
                    <h2>эксаба</h2>

                    <p>
                        это новое поколение анонимных досок,
                        проект призванный потеснить убогие кусабы и вакабы на анонимных просторах интернетов.
                        Доска запиливалась специально для анонимного общения с множеством пользовательских функций,
                        уникальная защита от вайпа, отсутсвие капчи, полная анонимность постинга, автодогрузка новых комментов в треде итд.
                        <a href="admin/">админка</a>
                    </p>
                    <div className="info">We are Anonymous. We are Legion. We do not forgive. We do not forget. Expect us.</div>
                </div>
            )
        } else {
            return (
                props?.page
            )
        }
    }
    return (
        <>
            <Menu />
            {content()}
            <Footer />
        </>
    );
}