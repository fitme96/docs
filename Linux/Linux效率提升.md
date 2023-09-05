

## 安装zsh, OhMyZsh

```bash
apt install -y zsh
# 切换至zsh
chsh -s $(which zsh)
# 安装ohmyzsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# 安装powerlevel0k 设置主题
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k

# 预测命令
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
# 命令检查
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
# vim ~/.zshrc
ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(git z sudo zsh-autosuggestions zsh-syntax-highlighting)

```




# 增加别名,输入文件名即可打开后缀文件
alias -s py=vim 
alias -s conf=vim
alias -s tgz='tar zxvf'
