import { useEffect } from 'react';

const useScript = (props: {url: string}) => {
    useEffect(() => {
        const script = document.createElement('script');

        script.src = props.url;
        script.async = true;

        document.body.appendChild(script);

        return () => {
            document.body.removeChild(script);
        }
    }, [props]);
};

export default useScript;