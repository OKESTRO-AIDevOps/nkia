import sys
import subprocess



PWD = sys.argv[1]

PWD_SOURCE = PWD + '/'

PWD += '/staging_downstream/'

REPO_LIST = []

REPO_ARG_DICT = {}


def RUN_MSG(message):

    print('* RUN_MSG: '+message)


def line_classifier(line):

    cls = ''

    hit = 0 

    for j in range(len(line)):

        if j == 0 and (line[j] != ' ' and line[j] != '-'):

            cls = 'REPO'
            hit = 1
            break

        elif j != 0 and line[j] == ' ':

            continue
        
        elif j != 0 and line[j] == '-':

            cls = 'ARG'
            hit = 1
            break

        elif j != 0 and (line[j] != ' ' and line[j] != '-'):

            cls = 'VAL'
            hit = 1
            break 

    if hit == 0 :
        cls = 'ILLEGAL'
        return

    return cls

def quick_sanitize(san_target):

    return san_target.replace(' ','').replace('\n','')


def merge_push_by_args(repo_address, git_dir, source_target, target_ignore):

    target_git_dir = PWD + git_dir + '/'


    _ = subprocess.run(['git','clone',repo_address])

    _ = subprocess.run(['mv', git_dir, PWD])



    for i in range(len(source_target)):

        source = source_target[i][0]

        target = source_target[i][1]

        source_effective = PWD_SOURCE + source

        target_effective = target_git_dir + target

        _ = subprocess.run(['rm','-rf',target_effective])

        _ = subprocess.run(['/bin/cp','-Rf',source_effective, target_effective])

    for i in range(len(target_ignore)):

        if target_ignore[i] == '':

            continue

        
        ignore = target_ignore[i]

        target_effective = target_git_dir + ignore

        _ = subprocess.run(['rm','-rf',target_effective])

    _ = subprocess.run(['git', '-C', target_git_dir ,'add','.'])

    _ = subprocess.run(['git', '-C', target_git_dir ,'commit','-m','"automated push by ./hack/feed-downstream"'])

    _ = subprocess.run(['git', '-C', target_git_dir ,'push'])


def cleanup():


    _ = subprocess.run(['rm','-rf', PWD +'nkia*'])




RUN_MSG('source root directory: '+ PWD_SOURCE)

RUN_MSG('target root directory: '+ PWD)


f = open(PWD+'repo.arg.fool')

raw_text = f.read()

f.close()

repo_arg_list = raw_text.split('\n')

try :

    repo_arg_list.remove('\n')

    repo_arg_list.remove('')

    RUN_MSG("null value detected and removed")

except:

    RUN_MSG("no null value detected")


CURRENT_REPO = ''

CURRENT_ARG = ''

for i in range(len(repo_arg_list)):
    
    cls = line_classifier(repo_arg_list[i])

    if cls == 'REPO':

        repo = repo_arg_list[i].split('=')

        CURRENT_REPO = quick_sanitize(repo[0])

        REPO_ARG_DICT[CURRENT_REPO] = {"REPO":quick_sanitize(repo[1]), "ARG":{}}

        REPO_LIST.append(CURRENT_REPO)

    elif cls == 'ARG':

        if CURRENT_REPO == '' :

            RUN_MSG('illegal behavior: ARG accessed before REPO')

            sys.exit('ABORT.')


        arg = repo_arg_list[i].split('=')

        CURRENT_ARG = quick_sanitize(arg[0].replace('-',''))


        if ',' in quick_sanitize(arg[1]):

            arg_list = arg[1].split(',')

            REPO_ARG_DICT[CURRENT_REPO]["ARG"][CURRENT_ARG] = [[quick_sanitize(arg_list[0]),quick_sanitize(arg_list[1])]]    

        else :

            REPO_ARG_DICT[CURRENT_REPO]["ARG"][CURRENT_ARG] = [quick_sanitize(arg[1])]
    
    elif cls == 'VAL':

        if CURRENT_ARG == '' :

            RUN_MSG('illegal behavior: VAL accessed before ARG')

            sys.exit('ABORT.')
        
        if ',' in repo_arg_list[i]:

            val = repo_arg_list[i].split(',')
    

            REPO_ARG_DICT[CURRENT_REPO]["ARG"][CURRENT_ARG].append([quick_sanitize(val[0]),quick_sanitize(val[1])])


        else :

            REPO_ARG_DICT[CURRENT_REPO]["ARG"][CURRENT_ARG].append(quick_sanitize(val))

    
    elif cls == 'ILLEGAL':

        RUN_MSG('illegal line detected near line: ' + str(i))

        sys.exit('ABORT.')



try: 
    for repo_key in REPO_LIST:

        repo_address = REPO_ARG_DICT[repo_key]["REPO"]

        git_dir = REPO_ARG_DICT[repo_key]["ARG"]["DOWNSTREAM_ROOT"][0]

        source_target = REPO_ARG_DICT[repo_key]["ARG"]["SOURCE_TARGET"]

        target_ignore = REPO_ARG_DICT[repo_key]["ARG"]["DOWNSTREAM_TARGET_IGNORE"]

        merge_push_by_args(repo_address, git_dir, source_target, target_ignore)


        
except:
    
    RUN_MSG('fatal error occurred')
    RUN_MSG('progress rolled back except for commits already made')
    RUN_MSG('ABORT.')


cleanup()

RUN_MSG('EXIT')





