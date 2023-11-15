import { FC } from "react";
import { useRouteError } from "react-router-dom";

const ErrorPage:FC = () => {
    const error: unknown = useRouteError();

    return (
        <div>
            <h1>Oops!</h1>
            <p>{(error as Error)?.message || (error as { statusText?: string })?.statusText}</p>
        </div>
    )
}

export default ErrorPage