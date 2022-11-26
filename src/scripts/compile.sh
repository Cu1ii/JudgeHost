# 创建提交工作目录
mkdir "$SUBMISSION_PATH";

# 创建代码路径，写入用户代码
touch "$CODE_PATH";

echo "$CODE" >> "$CODE_PATH";

BUILDING_SCRIPT_PATH="$SUBMISSION_PATH/build.sh"

# 创建编译脚本目录
touch "$BUILDING_SCRIPT_PATH";

# cd 到本次提交的工作目录
echo -e "$BUILDING_SCRIPT" >> "$BUILDING_SCRIPT_PATH";

chmod 777 "$BUILDING_SCRIPT_PATH";

# 执行编译
# shellcheck disable=SC2164
cd "$SUBMISSION_PATH"

# 初始化runner
touch run;

chmod 777 ./run;

echo "370802wsl" | sudo -S $JUDGE_CORE_PATH -t 4000 -c 4000 -m 100000 -f "$COMPILE_INFO_OUT_MAX_SIZE" -u "$USER_ID" -r "$BUILDING_SCRIPT_PATH" -o compile.out -e compile.err