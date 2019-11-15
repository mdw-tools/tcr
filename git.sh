function tcr_commit_count() {
    count=0
    echo $count
    git log --oneline | while read line ; do
        case "$line" in
            *tcr*)
                count=$((count+1));
            ;;
            *)
                return $count
        esac
    done
    return $count
}
