import { useCallback, useRef } from "react";
import { useInfiniteQuery } from "react-query";
import { Item } from "./Item";

const fetchUsers = async (ctx: { pageParam?: number }) => {
  const { pageParam = 1 } = ctx;
  const response = await fetch(`/apiURL/users?page=${pageParam}&per_page=50`);
  const data = await response.json();

  return data;
};

function UserList(): JSX.Element {
  const wrapperRef = useRef<HTMLDivElement>(null);
  const { data, error, fetchNextPage, hasNextPage, isFetchingNextPage } =
    useInfiniteQuery("users", fetchUsers, {
      getNextPageParam: (lastPage, allPages) => {
        const maxPages = lastPage.total_count / 30;
        const nextPage = allPages.length + 1;
        return nextPage <= maxPages ? nextPage : undefined;
      },
    });
  const intObserver = useRef<IntersectionObserver | null>(null);
  const lastUserNodeRef = useCallback(
    (userNode: HTMLLIElement) => {
      if (isFetchingNextPage) return;

      if (intObserver.current) intObserver.current.disconnect();

      intObserver.current = new IntersectionObserver((userNodes) => {
        if (userNodes[0].isIntersecting && hasNextPage) {
          fetchNextPage();
        }
      });

      if (userNode) intObserver.current.observe(userNode);
    },
    [isFetchingNextPage, fetchNextPage, hasNextPage]
  );

  if (error) return <p>{"An error has occurred: " + error}</p>;

  return (
    <div ref={wrapperRef}>
      {data?.pages.map((page) => {
        return page.userList.map(
          (user: { displayName: string; fid: number }, i: number) => {
            if (page.userList.length === i + 1) {
              return (
                <Item
                  ref={lastUserNodeRef}
                  key={user.fid}
                  text={user.displayName}
                  to={`/users/${user.fid}`}
                />
              );
            }
            return (
              <Item
                key={user.fid}
                text={user.displayName}
                to={`/users/${user.fid}`}
              />
            );
          }
        );
      })}
    </div>
  );
}

export default UserList;
