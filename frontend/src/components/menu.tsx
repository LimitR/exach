function Menu() {
    return (
        <div className="menu" id="header">

            <select onChange={(v)=> console.log(v.target.value)
            }>
                <option value="0" >Стили борды</option>
                <option value="none">Обычный</option>
                <option value="AtoX">AtoX</option>
                <option value="Neutron">Neutron</option>
            </select>
            <nav>
                <ul>
                    <li><a href="/">Глагне</a></li>
                    <li><a href="/galery" title="Галерея лучших изображений рунета и забугорья">Галерея</a></li>
                    <li><a href="/all" title="Все треды и посты в одном потоке">Однопоток</a></li>
                </ul>
            </nav>
        </div>
    )
}

export default Menu