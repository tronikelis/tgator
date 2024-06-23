import { RequireChildren } from "solid-daisy";

export default function Root(props: RequireChildren<Record<any, any>>) {
    return (
        <div class="w-full h-full flex items-center justify-center my-12">
            <div class="container">{props.children}</div>
        </div>
    );
}
